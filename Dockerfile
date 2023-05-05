FROM golang:1.20.4-bullseye

# Add make command
RUN apt update && apt install make

WORKDIR /app
COPY . /app
RUN make

WORKDIR /appbin
RUN cp /app/github-issue-schedule /appbin/
RUN rm -r /app

ENV PATH=${PATH}:/appbin

CMD ["github-issue-schedule"]
