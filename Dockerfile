FROM python:latest
WORKDIR /app/
COPY . /app/
RUN pip install toml
RUN pip install termcolor
CMD ["./jenrik", "./test_jenrik.toml"]
