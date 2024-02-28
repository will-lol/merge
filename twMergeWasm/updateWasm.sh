pushd lib
curl -L https://cdn.jsdelivr.net/npm/tailwind-merge@latest/dist/bundle-mjs.mjs > tw-merge.js
esbuild index.js --format=esm --bundle --outfile=bundle.js
javy compile bundle.js --wit index.wit -n index -o index.wasm
popd
