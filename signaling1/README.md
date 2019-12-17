# signaling1

*context.WithCancel* used for signaling a goroutine to stop, with the addition of
the goroutine reporting back its status via another *context.WithCancel*, 
waiting for the goroutine to have shut down before exiting.