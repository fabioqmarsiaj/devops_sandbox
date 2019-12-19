# Calculator made in GO and provisioned with GCP

A simple go calculator HTTP Service provided with Google Cloud Platform.

Google Cloud Platform it's a cloud computing suite offered by Google.
With Cloud Computing you can:

- Reduce IT costs: Rather than purchasing expensive systems you can use cloud resources, hardware as a software.
- Scalability: Offers flexibility as your needs change.
- Collaboration efficiency: If you are working on a project across different locations, you could use cloud computing to give employees, contractors and third parties access to the same files.

## Summary

- [Requirements](#requirements)
- [Installation](#installation)
- [Initializing GCP](#initializing-gcp)
- [Running the Application](#running-the-application)
- [Author](#author)

## Requirements

- [Git](https://git-scm.com/)
- [Google Cloud](https://cloud.google.com/?hl=pt-br)

## Installation

- Installing Git:

```
$apt-get install git
```

- Clone the project from github:

```
$git clone https://github.com/fabioqmarsiaj/jovens-talentos/tree/tema13
```

## Initializing GCP

You must have a Google Cloud Account, to do so create one at: https://cloud.google.com/

- Access the project folder ~/jovens-talentos/3-devops/fabio-marsiaj/Tema14/ and run:

```
$curl https://sdk.cloud.google.com | bash
```

- Now to log in your account and initialize the GCP in your project do as following:

```
$exec -l $SHELL
$sudo chown {your-pc-user} -R /home/{your-pc-user}/.config/gcloud
$gcloud init
```

- You are able to create a GCP project. You could do that on the GCP Console as well, but with the terminal run this:

```
$gcloud projects create {projectName}
```

## Running the Application

Now you are good to go!

To deploy your applicationg to GCP run:

```
$gcloud app deploy
```

And done! You can check the go calculator by running:

```
$gcloud app browse
```

Endpoints:

You can make simple calculations such as: sum, sub, mul or div by accessing the following endpoints:

https://{projectName}.appspot.com/calc/{operation}/{number1}/{number2}

At {operation} choose: sum, sub, mul or div.

If you want to check out the calculator history just access:
http://{projectName}.appspot.com/calc/history

## Author

**Fabio Quinto Marsiaj** - [GitHub](https://github.com/fabioqmarsiaj)

<a href="https://github.com/fabioqmarsiaj">
    <img
    alt="Imagem do Autor Fabio Marsiaj" src="https://avatars0.githubusercontent.com/u/34289167?s=460&v=4" width="100">
</a>
