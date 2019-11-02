FROM alpine
ADD example2-srv /example2-srv
ENTRYPOINT [ "/example2-srv" ]
