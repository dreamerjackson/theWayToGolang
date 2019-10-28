package mapreduce

import (
	"fmt"
	"sync"
	"log"
)

//
// schedule() starts and waits for all tasks in the given phase (Map
// or Reduce). the mapFiles argument holds the names of the files that
// are the inputs to the map phase, one per map task. nReduce is the
// number of reduce tasks. the registerChan argument yields a stream
// of registered workers; each item is the worker's RPC address,
// suitable for passing to call(). registerChan will yield all
// existing registered workers (if any) and new ones as they register.
//
func schedule(jobName string, mapFiles []string, nReduce int, phase jobPhase, registerChan chan string) {
	var ntasks int
	var n_other int // number of inputs (for reduce) or outputs (for map)
	switch phase {
	case mapPhase:
		ntasks = len(mapFiles)
		n_other = nReduce
	case reducePhase:
		ntasks = nReduce
		n_other = len(mapFiles)
	}

	fmt.Printf("Schedule: %v %v tasks (%d I/Os)\n", ntasks, phase, n_other)

	// All ntasks tasks have to be scheduled on workers, and only once all of
	// them have been completed successfully should the function return.
	// Remember that workers may fail, and that any given worker may finish
	// multiple tasks.

	// schedule will wait until all worker has done their jobs
	var wg sync.WaitGroup

	// RPC call parameter
	var task DoTaskArgs
	task.JobName = jobName
	task.NumOtherPhase = n_other
	task.Phase = phase

	// task id will get from this channel
	var taskChan = make(chan int)
	go func() {
		for i := 0; i < ntasks; i++ {
			wg.Add(1)
			taskChan <- i
		}
		// wait all workers have done their job, then close taskChan
		wg.Wait()
		close(taskChan)
	}()

	// assign all task to worker
	for i := range taskChan {
		// get a worker from register channel
		worker := <-registerChan

		task.TaskNumber = i
		if phase == mapPhase {
			task.File = mapFiles[i]
		}

		// Note: must use parameter
		go func(worker string, task DoTaskArgs) {
			if call(worker, "Worker.DoTask", &task, nil) {
				// only successful call will call wg.Done()
				wg.Done()

				// put idle worker back to register channel
				registerChan <- worker;
			} else {
				log.Printf("Schedule: assign %s task %v to %s failed", phase,
					task.TaskNumber, worker)

				// put failed task back to task channel
				taskChan <- task.TaskNumber
			}
		}(worker, task)
	}
	fmt.Printf("Schedule: %v phase done\n", phase)
}
