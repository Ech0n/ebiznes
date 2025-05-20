FROM node:18

WORKDIR /client

COPY . .

RUN npm ci

EXPOSE 5173 5173

CMD ["npm", "run", "dev", "--host"]