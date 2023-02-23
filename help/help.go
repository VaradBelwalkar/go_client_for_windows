package help
 
import (
       "fmt"
)
//Just used to export information
var help = `
NAME
       dyplug : Client to access microservices

dyplug COMMANDS

       dyplug > login
              
              Login to access the services

       dyplug > config

               Displays the configured URL, username, and password of the client

       dyplug > change config

              Edits configuration and saves the changes
	   

       dyplug > upload file <file_path> <container_path> <container_name>

              uploads the specified file into the cloud

       dyplug > upload folder <dir_path> <container_path> <container_name>

              uploads the specified directory to the cloud

       dyplug > download file <filepath_in_container> <local_path> <container_name> 

              downloads the specified file from the cloud
       
       dyplug > download folder <folderpath_in_container> <local_path> <container_name>


dyplug COMMANDS TO ACCESS RUNTIMES

  (CURRENTLY AVAILABLE RUNTIMES : ubuntu , development_server)

  (You can own maximum of 5 containers at a time)


       dyplug > container run <runtime_name> <port_to_expose(optional)>
             
              get a new specifed container

       dyplug > container list images

              get available os images

       dyplug > container list containers

              get the list of containers that you own

       dyplug > container start <container_name>

              get the specified container that you own from your owned_containers list

       dyplug > container stop <container_name>

              stop the specified container on server side

       dyplug > container remove <container_name>

              remove the specified container from the server


`


func Help(){
       fmt.Println(help)
}
