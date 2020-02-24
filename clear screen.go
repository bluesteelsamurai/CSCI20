/*
func Clearscreen(){//clears screen using native language of operating screen
  cmd := exec.Command("cls") //inputs command directly to system --useful for other functions!!!
  system :=runtime.GOOS  //identifies what the system is that the program runs on...needed for above!
  if system== "linux"{//catches linux os's
    exec.Command("clear")
  }else if system == "darwin"{//catches mac os's
    exec.Command("clear")
  }
  cmd.Stdout()=os.Stdout()
  cmd.Run()
}*/
