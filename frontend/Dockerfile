# Step 1: Use an official Node runtime as a parent image to build the app
FROM node:20-alpine AS builder
WORKDIR /app

# Copy lockfiles and install dependencies
COPY package*.json ./
RUN npm install

# Copy your source code and compile the static production assets
COPY . .
RUN npm run build

# Step 2: Use a lightweight Nginx web server to host the static assets fast
FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
