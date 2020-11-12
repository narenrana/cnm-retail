# pull official base image
FROM node:13.12.0-alpine

# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
COPY ./payment-app/package.json ./
COPY ./payment-app/public ./public
COPY ./payment-app/src ./src
COPY ./payment-app/server.js ./

RUN ls ./

RUN npm install
RUN npm run build

RUN ls ./
# add app

RUN ls
RUN pwd

# start app
CMD ["node", "server.js"]