FROM node:20.11.0-alpine

WORKDIR /home/app

COPY package.json ./
COPY tsconfig.json ./
COPY tsconfig.build.json ./
COPY src ./src

RUN npm install

EXPOSE 5000

CMD ["npm", "run", "start:dev"]
