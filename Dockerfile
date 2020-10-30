FROM python:3.7-slim

ENV M_WORKDIR "/workdir"
ENV M_RESOURCES "/resources"
ENV M_SHARED "/shared"

COPY resources /resources
COPY workdir /workdir

WORKDIR /workdir

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y git make curl wget &&\
    wget https://github.com/mikefarah/yq/releases/download/3.3.4/yq_linux_amd64 -O /usr/bin/yq &&\
    chmod +x /usr/bin/yq 

RUN git clone https://github.com/kubernetes-sigs/kubespray.git

RUN pip3 install -r ./kubespray/requirements.txt


# to be removed

#RUN cp -rfp inventory/sample inventory/mycluster
#RUN cp -rfp /resources/inventory.ini inventory/mycluster/


#RUN cat inventory/mycluster/group_vars/all/all.yml
#RUN cat inventory/mycluster/group_vars/k8s-cluster/k8s-cluster.yml
# ...

ENTRYPOINT ["make"]
