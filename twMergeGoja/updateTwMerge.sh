pushd lib
curl -L https://cdn.jsdelivr.net/npm/tailwind-merge@latest/dist/es5/bundle-mjs.mjs > tw-merge.js
esbuild tw-merge.js --target=es5 --bundle --outfile=bundle.js --global-name=m
popd
