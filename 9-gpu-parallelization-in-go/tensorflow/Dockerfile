FROM tensorflow/tensorflow:latest-gpu
ENV FOO=BAR
ENV LD_LIBRARY_PATH=/tmp:$LD_LIBRARY_PATH
RUN ln -s /usr/local/cuda-10.1/lib64/libcudart.so /usr/lib/libcudart.so
RUN add-apt-repository ppa:longsleep/golang-backports
RUN apt update
RUN apt -y install golang-go
COPY . /src
WORKDIR /src
RUN apt install wget
RUN wget https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-gpu-linux-x86_64-1.15.0.tar.gz
RUN tar -C /usr/local -xzf libtensorflow-gpu-linux-x86_64-1.15.0.tar.gz
RUN ldconfig
#RUN apt -y install nvidia-utils-440 nvidia-modprobe
#RUN pip install tensorflow

# SYMLINK LIBRARIES
# INSTALL CUDNN

ENTRYPOINT ["/usr/bin/go", "run", "tf.go", "2&>/dev/null"]
