## Setup
Before running create *config.yml* by *config.example.yml* in /config folder. <br>
To run the application use the command below.
> make startapp


## Migrations
Before using migrations [install](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md) *golang-migrate*. <br>

Use the command below to up the migrations.
> migrate -path db/migrations -database postgresql://postgres:111111@localhost:5432/ads?sslmode=disable up

Or use the command below to down the migrations.
> migrate -path db/migrations -database postgresql://postgres:111111@localhost:5432/ads?sslmode=disable down


## Amazon S3 configuration
To configure create directory .aws in %UserProfile% folder in Windows and $HOME or ~ (tilde) in Unix-based systems.

1. Create credentials file in .aws directory
> [default] <br>
> aws_access_key_id=YOUR_ACCESS_KEY_ID <br>
> aws_secret_access_key=YOUR_SECRET_ACCESS_KEY

Check how to [get access key ID and secret access key](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html#cli-configure-quickstart-creds). 

2. Create config file in .aws directory
> [default] <br>
> region=eu-central-1

Check your nearest [regional endpoints](https://docs.aws.amazon.com/general/latest/gr/rande.html).

## MailGun
To configure mailgun, [sign up or log in](https://www.mailgun.com). Then [find private API key](https://help.mailgun.com/hc/en-us/articles/203380100-Where-Can-I-Find-My-API-Key-and-SMTP-Credentials-) and [domain name](https://beta.mailgun.com/mg/dashboard).
Copy and paste private API key and domain name to *config.yml* <br>
For testing [add an authorized recipient](https://help.mailgun.com/hc/en-us/articles/217531258-Authorized-Recipients).


## Routes
To use all routes download *Goads Collection.postman_collection.json* from repository and import this file to [Postman](https://www.postman.com).
