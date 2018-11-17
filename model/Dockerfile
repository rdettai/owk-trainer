FROM python:2.7-slim

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install any needed packages
RUN pip install keras
RUN pip install tensorflow

# Run app.py when the container launches
CMD ["python", "algo.py"]