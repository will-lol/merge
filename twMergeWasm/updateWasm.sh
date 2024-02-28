pushd lib
curl -L https://cdn.jsdelivr.net/npm/tailwind-merge@latest/dist/bundle-mjs.mjs > tw-merge.js
esbuild index.js --bundle --outfile=bundle.js
javy compile bundle.js -o index.wasm
popd
