FROM node:22-alpine

RUN mkdir -p /home/node/app/node_modules && chown node:node /home/node/app
WORKDIR /home/node/app
COPY package*.json ./

RUN npm ci && \
  npm config --global set update-notifier false && \
  npm cache clean --force
COPY . .

USER node
ENTRYPOINT ["npm", "run"]
