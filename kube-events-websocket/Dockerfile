FROM python:3.10-alpine
WORKDIR /app
COPY . /app
RUN pip install -r requirements.txt
ENTRYPOINT ["flask", "run", "--host=0.0.0.0", "--port=80"]