FROM alpine:3.11
ADD hmacCalcTools /hmacCalcTools
ADD tls /tls
ENTRYPOINT ["/hmacCalcTools"]