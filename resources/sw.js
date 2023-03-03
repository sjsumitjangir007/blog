
const cacheName = "js13kPWA-v1";
const appShellFiles = [
  "/",
  "/posts",
  "/post/why_weblocks_special",
  "/post/ngrok_kickstart",
  "/post/golang_optional_values",
  "/post/democratic_nature_of_scrum",
  "/assets/sms_icon_favicon_transparent.svg",
  "/assets/me_raw.jpg",
  "/assets/core.js",
  "/favicon.ico",
];

self.addEventListener("install", (e) => {
    console.log("[Service Worker] Install");
    e.waitUntil(
      (async () => {
        const cache = await caches.open(cacheName);
        console.log("[Service Worker] Caching all: app shell and content");
        await cache.addAll(appShellFiles);
      })()
    );
  });

  self.addEventListener("fetch", (e) => {
    e.respondWith(
        (async () => {
          const r = await caches.match(e.request);
          console.log(`[Service Worker] Fetching resource: ${e.request.url}`);
          if (r) {
            return r;
          }
          const response = await fetch(e.request);
          const cache = await caches.open(cacheName);
          console.log(`[Service Worker] Caching new resource: ${e.request.url}`);
          cache.put(e.request, response.clone());
          return response;
        })()
      );
  });

  self.addEventListener("activate", (e) => {
    e.waitUntil(
      caches.keys().then((keyList) => {
        return Promise.all(
          keyList.map((key) => {
            if (key === cacheName) {
              return;
            }
            return caches.delete(key);
          })
        );
      })
    );
  });