FROM python:3.10-alpine
WORKDIR /app
COPY . /app
EXPOSE 8080
RUN pip install -r requirements.txt
ENTRYPOINT ["python3","server.py"]