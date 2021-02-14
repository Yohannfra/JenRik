FROM python:latest
WORKDIR /app/
COPY . /app/
RUN pip3 install -r ./requirements.txt
CMD ["./jenrik", "./test_jenrik.toml"]
