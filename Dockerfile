FROM docker.pkg.github.com/bryk-io/base-images/shell:0.1.0

# Expose required ports
EXPOSE 9090

# Expose required volumes
VOLUME /etc/affinityctl

# Add application binary and use it as default entrypoint
COPY affinityctl_linux_amd64 /bin/affinityctl
ENTRYPOINT ["/bin/affinityctl"]
