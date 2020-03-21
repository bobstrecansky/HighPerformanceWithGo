FROM tensorflow/tensorflow:latest-gpu
ENV LD_LIBRARY_PATH=/usr/local/cuda-10.1/lib64
RUN ln -s /usr/local/cuda-10.1/lib64/libcudart.so /usr/lib/libcudart.so
RUN apt-get install -y golang
COPY . /tmp
WORKDIR /tmp
RUN make
RUN mv libmultiply.so /usr/lib/libmultiply.so          
ENTRYPOINT ["/usr/bin/go", "run", "cuda_multiply.go"]   
