FROM public.ecr.aws/lambda/provided:al2
COPY ./build/kaytu-azure-describer ./
ENTRYPOINT [ "./kaytu-azure-describer" ]
CMD [ "./kaytu-azure-describer" ]