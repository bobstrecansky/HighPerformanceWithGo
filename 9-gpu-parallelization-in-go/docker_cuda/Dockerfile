FROM tensorflow/tensorflow:latest-gpu
ENV LD_LIBRARY_PATH=/tmp:$LD_LIBRARY_PATH
RUN ln -s /usr/local/cuda-10.1/lib64/libcudart.so /usr/lib/libcudart.so
RUN apt-get install -y golang
COPY . /tmp
WORKDIR /tmp
RUN make
ENTRYPOINT ["/usr/bin/go", "run", "cuda_multiply.go"]
