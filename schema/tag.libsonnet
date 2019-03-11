local tagExample = {
  name: 'v1.0.1',
  clients: {
    exampleCli: {
      download: [
        'example-cli',
        'download',
        '--image',
        'repository-foo',
        '--tag',
        'v1.0.1',
      ],
      upload: [
        'example-cli',
        'push',
        '--image',
        'repository-foo',
        '--tag',
        'v1.0.1',
      ],
    },
  },
  mediaType: 'application/vnd.oci.image.manifest.v1+json',
  digest: 'sha256:10b995d6204131069af3e4f00dc1d3758d517a5edb29e5757d3c2858d5613127',
  size: 435,
};

{
  tag:: {
    example:: tagExample,
  },
}