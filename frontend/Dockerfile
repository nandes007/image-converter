FROM node:21-alpine

WORKDIR /usr/src/fronted
COPY . .
RUN npm install -g @angular/cli @angular-devkit/build-angular && npm install --force
RUN npm install

CMD ["ng", "serve", "--host", "0.0.0.0", "--disable-host-check"]

