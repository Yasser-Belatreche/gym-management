# Build Stage 1
# This build created a staging docker image
#
FROM node:20.11.0-alpine as StagingBuild

WORKDIR /home/app

COPY package.json ./
COPY tsconfig.json ./
COPY tsconfig.build.json ./
COPY src ./src

RUN npm install

RUN npm run build

# Build Stage 2
# This build takes the production build from staging build
#
FROM node:20.11.0-alpine as ProductionBuild

WORKDIR /home/app

COPY package.json ./
COPY --from=StagingBuild /home/app/dist ./dist


RUN npm install --omit=dev
RUN npm install -g pm2

RUN chown -R node:node /home/app

USER node

ENTRYPOINT ["pm2-runtime", "dist/index.js"]
