$file_or_folder=$args[0]
$path=$args[1]
$containerpath=$args[2]
$port=$args[3]
$server=$args[4]
$keyForRemoteServer=$args[5]

if (${file_or_folder} -eq "file"){
  scp -i ${keyForRemoteServer} -P ${port} ${path} root@${server}:${containerpath}}
elseif(${file_or_folder} -eq "folder"){
  scp -r -i ${keyForRemoteServer} -P ${port} ${path} root@${server}:${containerpath}}
else{
  Write-Host "Error: Invalid argument. Expecting either 'file' or 'folder'."
}