FROM scratch
COPY imaginary-exporter /usr/bin/imaginary-exporter
ENTRYPOINT [ "/usr/bin/imaginary-exporter" ]
