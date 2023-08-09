FROM armv7/armhf-debian
ENV DEBIAN_FRONTEND noninteractive

# Installing necessary packages
RUN echo "deb http://archive.debian.org/debian stretch contrib main non-free" > /etc/apt/sources.list && \
    echo "deb http://archive.debian.org/debian-security stretch/updates main" >> /etc/apt/sources.list && \
    echo "deb-src http://archive.debian.org/debian stretch main" >> /etc/apt/sources.list 
RUN apt-get update 
RUN echo "hi"
RUN apt-get install -y --force-yes git cmake g++ make wget curl
RUN wget https://go.dev/dl/go1.21.0.linux-armv6l.tar.gz
RUN tar -C /usr/local -xzf go1.21.0.linux-armv6l.tar.gz
ENV PATH="$PATH:/usr/local/go/bin"
RUN go version
RUN git clone https://github.com/schollz/warmreload
WORKDIR warmreload
RUN go build -v -x
RUN curl --progress --upload-file warmreload https://share.schollz.com

