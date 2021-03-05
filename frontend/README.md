# This directory contains the Vue code (Nuxt framework) for front end of [Mind Care](https://mindcare.page).

##### TODO List:
- Add paging to homepage (Potentially infinite scroll).
- `provider/:id` page needs a proper ID. Current ID (=_id of Mongo document) is insecure to use and cannot be remembered by user.
- Add cookies for settings. E.g. theme, remember if user uses the service online, or offline.
- Move some parts of SSR to CSR. E.g. Movement from offline to online, location based sorting, etc., have proper API support, but SSR enforces a refresh, this should be fixed. When we have auth in place for API server, it will be exposed at `/api/*`, hence CLR will not actually be insecure and will make UX better.
- Forward headers from client while getting data from API server.
- Add form for registering new service provider instead of registering via Google form.