FROM python:3.7-slim

COPY resources /resources
WORKDIR /workdir

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y git

RUN git clone https://github.com/kubernetes-sigs/kubespray.git

RUN pip3 install -r ./kubespray/requirements.txt

ENTRYPOINT ["bash"]
