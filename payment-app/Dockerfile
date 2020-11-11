# pull official base image
FROM node:13.12.0-alpine

# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
COPY ./payment-app/package.json ./
 
RUN npm install --silent
RUN npm install react-scripts@3.4.1 -g --silent

RUN ls ./
# add app
COPY ./payment-app ./
RUN ls
RUN pwd

# start app
CMD ["npm", "start"]