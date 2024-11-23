package main
import ("net"; "fmt"; "flag"; "strconv")


// this will need to run multiple "Dial" functions for each port
func Port_Scan (network_connection string, IP_ports string) (net.Conn,error){
  // IP ports will be an array of strings in the future 
  
  conn , err := net.Dial(network_connection , IP_ports)
  return conn , err


}




func main (){
  var single_port int
  var Ip string
  var network_connection string 
  var IP_port_single string 


  network_connection = "tcp"
  
  // added flag syntax to be more professional
  
  flag.IntVar(&single_port, "Ss", -1, "Ss is used to scan a single port")
  flag.StringVar(&Ip, "Ip", "golang.org", "Ip is where you can enter the address of the server you would like to scan")


  // Make sure no flags appear after this line######################################################################### 
  //###################################################################################################################
  flag.Parse()





  // this is necessary because of how the "Dial" function works
  // in the future this will need to be its own function that uses a for loop to cover all the ports
  if single_port != -1{
    string_single_port := strconv.Itoa(single_port)
    IP_port_single = Ip + ":" + string_single_port
//    fmt.Println("IP_port_single was declared")
//    fmt.Println(IP_port_single)
  }else {
    IP_port_single = ""
    _ = IP_port_single 

  }
  conn , err := Port_Scan( network_connection, IP_port_single)
  if err != nil {
    fmt.Println("You had an error")
    fmt.Println(err)
  }
  if conn != nil {
  fmt.Println("Successfully connected to:", conn.RemoteAddr())
}
  




}// end of main function
