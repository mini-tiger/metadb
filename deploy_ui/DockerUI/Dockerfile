FROM node:16.14-buster

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm config set registry https://registry.npm.taobao.org &&   npm install vue-cli && npm install

COPY . .

EXPOSE 8080
CMD [ "npm", "run", "serve" ]