#!/usr/bin/perl

sub checkerr {
	if (@_[0]){
		print("We Ran into an error while doing")
	}
	exit();

}


sub installSnoopy {

	`mv /tmp/scripts/ld.so.preload /etc/ld.so.preload`;
	`mv /tmp/scripts/libsnoopy.so /usr/local/lib/libsnoopy.so`;
	`mv /tmp/scripts/snoopy.ini /etc/snoopy.ini`;
	
	# Restarting Services for better ld to be preloaded
	
	print("Restarting Apache2 Webserver...\n");

	#`systemctl restart apache2 2>/dev/null`;
	
	#checkerr($?);

	print("Success! logging utility was sucessfuly installed.");
	print("You can view your logs at /var/log/auth.log\n");
}

sub checkforweakSudo {
	$ouput = `cat /etc/sudoers | grep -v '#' | grep NOPASSWD`;
	if ($output!=""){
		print("Weak sudo permission found.\n");
		print("It is never adivisable to give user NOPASSWD on sudo\n");
		print($output);
		print("Check here to fix it:-");
		print("INFO: https://serverfault.com/questions/615034/disable-nopasswd-sudo-access-for-ubuntu-user-on-an-ec2-instance\n");
	}
}


# installSnoopy();
checkforweakSudo();