# build environment
FROM node:latest AS builder

WORKDIR /app

COPY package.json .
COPY package-lock.json .

RUN npm install --silent

COPY . .

RUN npm run build


# production environment
# running nginx with non-root user
FROM nginxinc/nginx-unprivileged

COPY --from=builder /app/dist /usr/share/nginx/html

# Some hosting service requires to listen on different port than 8080. In this case, there should be some twist to update the nginx config.
EXPOSE 8080

CMD ["nginx", "-g", "daemon off;"]




