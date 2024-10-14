FROM public.ecr.aws/lambda/provided:al2
COPY ./build/og-azure-describer ./
ENTRYPOINT [ "./og-azure-describer" ]
CMD [ "./og-azure-describer" ]