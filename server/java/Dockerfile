FROM amazoncorretto:17-alpine AS build

WORKDIR /src
COPY src .
RUN set -ex \
  && ls -al \
  && javac -d ./bin com/company/Main.java com/company/Server.java \
  && ls -al bin/com/company/

CMD ["sh", "-c", "time java -classpath ./bin com.company.Main"]
