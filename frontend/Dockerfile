FROM node:12-alpine
WORKDIR /mind-care
COPY package.json .
RUN npm install
COPY . .
RUN npm run build

CMD ["npm", "start"]
