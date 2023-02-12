$file_or_folder=$args[0]
$containerpath=$args[1]
$localpath=$args[2]
$port=$args[3]
$server=$args[4]
$keyForRemoteServer=$args[5]
if ($file_or_folder -eq "file"){
  scp -i ${keyForRemoteServer} -P ${port}  root@${server}:${containerpath} ${localpath}}

elseif (${file_or_folder} -eq "folder"){
  scp -r -i ${keyForRemoteServer} -P ${port}  root@${server}:${containerpath} ${localpath}}
else{
  Write-Host "Error: Invalid argument. Expecting either 'file' or 'folder'."}

