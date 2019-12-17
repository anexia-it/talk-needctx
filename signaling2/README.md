# signaling2

*context.WithCancel* used for signaling a goroutine to stop, with the addition of
the goroutine reporting back its status via another *context.WithCancel* and a timeout 
waiting for the goroutine to stop during shutdown.