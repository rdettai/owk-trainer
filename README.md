# Recruitment test OWKIN

## Step 1: build the python model docker image

In order to be able to test the API, we build the python script into a docker image. Also I am on windows and installing keras/tenserflow python dep is painfull so we might as well use docker directly! The Docker file can be found in the model/ folder.

No effort was given on ensuring that the Keras/Tensorflow are using the GPU. The training seems pretty slow so my guess would be that it didn't.

## Step 2: share the files

We have to mount a volume to share the data between models and the API service
- docker run -v ModelVolume:/model rdettai/owkin-model-1:rev1

I had some difficulties checking that the model could actually write to the folder it was told to in the Dockerfile because the model took a long time to execute. I should have tested it with a mock model. 

I would prefer to have the models send their results to a file storage server (in the cloud that would be S3) or a database. This could be achieved by adding a side-car container that forwards everything that is written to the file system.

The image can be pushed to Dockerhub, for instance as rdettai/owkin-model-1:rev1

## Step 3: the GET /v1/models score api

The first objective is to list the shared volume to list existing scores. For simplicity, we made an api that lists all score instead of an API that lists individual scores. The name of the model is simply the folder in which the score was written, this could also be improved. Note that for testing, we have added a test/ folder with the same structure as the one we want in the ModelVolume shared volume.

## Step 3: the POST /v1/models model api

The api will take a docker image, connect to the docker socket and launch a container that is bound to the ModelVolume. 

This part is a little bit more tricky. It is kind of complicated to test this on windows because I don't know how to connect to the docker socket, and I don't have the time to setup a VM to do this.

I just copied some sample code from a library that is supposed to act as a docker client, but as I didn't manage to have the connection I didn't adapt it to start new containers.

The input object for the method POST /v1/models would be
{
  "image": "registry.hub.docker.com/rdettai/owkin-model-1:rev1",
  "name": "mnist-keras"
}

The name could then be passed as an env variable to the image to determine the folder in thich the score is written, this would imply to force the model container authors to have the folder in which they write as an env varibale which is maybe not ideal.

## Step 4: build the go api docker image

Now we create a docker image for the api. The Dockerfile is directly in the server/ folder. The build first failed because of the Go docker client library, probably because I was using an alpine build image. I removed the -alpine but I can't test it because docker-machine just crashed again.
- docker run -v ModelVolume:/model -v /var/run/docker.sock:/var/run/docker.sock -p 8081:8081 rdettai/owkin-model-1:rev1




## Conclusion
- I feel that I am not very far from having something that clicks in and works
- I had some difficulties building docker images because I use docker-machine which is slow and tends to crash when run multiple build commands simulteanously
- I am only half confident that making the connection to the docker socket would work like magic, but I remember having already done something similar.
- The design for sharing information through a shared volume is not very scalable and I doubt its stability. I would realy prefer to do it in some kind of object database.
- Security was not taken into account.
- I took a bit more than 3h because I didn't want to stop! But I really need to go... ;-)



## Development
For development, we recommand using https://github.com/codegangsta/gin for hot reload
- cd server
- gin -i run

