FROM ubuntu:16.04
RUN apt-get update -y && \
    apt-get install -y git curl python build-essential
RUN git clone https://github.com/IBM/monitoring_ui
RUN cd monitoring_ui && ./setup.sh
EXPOSE 8081
WORKDIR monitoring_ui
CMD /bin/bash -c ". ~/.bash_profile && npm run dev-server"
