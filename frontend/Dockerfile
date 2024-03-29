FROM node:21-alpine AS build
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm cache clean --force
RUN npm install
RUN npx ngcc --properties es2023 browser module main --first-only --create-ivy-entry-points
COPY . .
RUN npm run build
FROM nginx:stable
COPY --from=build /usr/src/app/dist/frontend/browser /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf
