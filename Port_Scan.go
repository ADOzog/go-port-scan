package main
import ("net"; "fmt"; "flag"; "strconv"; "strings"; "time"; "sync")


// this will need to run multiple "Dial" functions for each port
func Port_Scan (network_connection string, IP_ports string, timeout_duration time.Duration, wg *sync.WaitGroup) {
  // IP ports will be an array of strings in the future 
  defer wg.Done() 
  conn , err := net.DialTimeout(network_connection , IP_ports, timeout_duration)
    _ = err
    // the error handleing got a bit complex but in the future the goal is to make i/o timeout go away or be made simple
    // if err != nil {
    // if netErr , ok := err.(net.Error); {
    //     fmt.Println("i/o")
    //   }else{    
    //     fmt.Println("You had an error")
    //     fmt.Println(err)
    //   }
    // }
    if err != nil {
    fmt.Println(err)
  }  
    if conn != nil {
      fmt.Println("Successfully connected to:", conn.RemoteAddr())
    }


}

func IP_port_maker(multiple_ports_array []int,IP string)([]string){
  array_length := len(multiple_ports_array)
  IP_port_array := make([]string, array_length)

  for i := range(array_length){
    IP_port_array[i] = IP + ":" + strconv.Itoa(multiple_ports_array[i])

  }
  return  IP_port_array
}// this makes the array of the ip combined with the port




func main (){
  // var single_port int
  var multiple_ports_string_form string
  var multiple_ports_array []int
  var timeout_int int
  // var set_of_ports [...]int
  

  var IP string
  var network_connection string 
  // var IP_port_single string 
  
  var wg sync.WaitGroup

  network_connection = "tcp"
  
  // added flag syntax to be more professional
  
  // flag.IntVar(&single_port, "Ss", -1, "Ss is used to scan a single port")
  flag.StringVar(&IP, "IP", "golang.org", "Ip is where you can enter the address of the server you would like to scan")

  flag.StringVar(&multiple_ports_string_form, "Sm", "", "Sm is used to scan multiple ports, just write each port you would like scan seperated by spaces and surrounded by '' ")

  flag.IntVar(&timeout_int, "T", 5, "T is used to set the amout of time till a port is considered failed to connect")


  Scan_all := flag.Bool( "Sa", false, "Sa is used to scan all ports from 1 to 1024")

  // Make sure no flags appear after this line######################################################################### 
  //###################################################################################################################
  flag.Parse()


  timeout_duration := time.Duration(timeout_int) * time.Second


  // doing this with slices may impact the speed I will look at this in the future
  if *Scan_all {
    for i := range(1024){
      multiple_ports_array = append(multiple_ports_array,i)
    }
  }



  multiple_ports_array_of_strings := strings.Fields(multiple_ports_string_form)
  for _,port_string := range multiple_ports_array_of_strings{
    num,err := strconv.Atoi(port_string)
    if err != nil {
      fmt.Println("Error converting input from '-Sm' to number ")
      continue
    }
    multiple_ports_array = append(multiple_ports_array, num)
  }
  // this is necessary because of how the "Dial" function works
  // in the future this will need to be its own function that uses a for loop to cover all the ports
   
  Array_IP_ports := IP_port_maker(multiple_ports_array , IP)
  
  for _,IP_port := range Array_IP_ports{ 
    wg.Add(1)
    go Port_Scan( network_connection, IP_port, timeout_duration, &wg)


  }
  wg.Wait()
 



}// end of main function
