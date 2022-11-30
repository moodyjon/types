import os
from setuptools import setup, find_packages

BASE = os.path.dirname(__file__)
with open(os.path.join(BASE, 'README.md'), encoding='utf-8') as fh:
    long_description = fh.read()

setup(
    name="lbry-types",
    version=2.3,
    author="LBRY Inc.",
    author_email="hello@lbry.com",
    url="https://lbry.com",
    description="Cross-language definitions for standard LBRY types.",
    long_description=long_description,
    long_description_content_type="text/markdown",
    keywords="lbry protocol media",
    license='MIT',
    python_requires='>=3.7',
    package_dir={
        'lbry_types': './python',
    },
    packages=[
        'lbry_types.v1',
        'lbry_types.v2',
    ],
    zip_safe=False,
    install_requires=[
        'protobuf==3.20.1',
    ],
)