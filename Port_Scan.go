package main
import ("net"; "fmt")


// this will need to run multiple "Dial" functions for each port
func Scan_Ports (network_connection string, IP_ports string) (net.Conn,error){
  // IP ports will be an array of strings in the future 
  
  conn , err := net.Dial(network_connection , IP_ports)
  return conn , err


}




func main (){
  var ports string
  var IP string
  var network_connection string 

  network_connection = "tcp"
  ports  = "80"
  fmt.Println("what network address?")
  //fmt.Scan(&IP)
  IP = "golang.org"

  // this is necessary because of how the "Dial" function works
  // in the future this will need to be its own function that uses a for loop to cover all the ports
  IP_ports := IP + ":" + ports

  conn , err := Scan_Ports( network_connection, IP_ports)
  if err != nil {
    fmt.Println("You had an error")
    fmt.Println(err)
  }
  if conn != nil {
  fmt.Println("Successfully connected to:", conn.RemoteAddr())
}
  




}// end of main function
