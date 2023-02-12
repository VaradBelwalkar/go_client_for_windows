$port = $args[0]
$ipAddress = $args[1]
$privateKeyPath = $args[2]
ssh -i $privateKeyPath -p $port root@$ipAddress