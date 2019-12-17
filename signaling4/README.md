# signaling4

*context.WithCancel* used for signaling a goroutine to stop, using a *sync.WaitGroup* to report back its status.

This shows how to recover from panics inside goroutines and can, with a slight modification of removing the recover 
call, be used to demonstrate that a goroutine that panics takes down the entire process.
