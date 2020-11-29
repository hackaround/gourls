# gourls

URL shortener in Go.

## Development

### Create Your Databases

Ok, if you don't edit the "database.yml" file and started your database, now Buffalo can create the databases in that file for you:

    $ docker run -d --name gourls-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root mysql:8.0
	$ buffalo pop create -a

## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.

**Congratulations!** You now have your Buffalo application up and running.

## What Next?

We recommend you heading over to [http://gobuffalo.io](http://gobuffalo.io) and reviewing all of the great documentation there.

Good luck!

[Powered by Buffalo](http://gobuffalo.io)
