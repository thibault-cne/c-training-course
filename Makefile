# You need to specify every Exercise folders here in  order to build them.
FOLDERS = Exercise1 Exercise2 Exercise3

SOLUTIONS = solution/
BASEPATH = ..

all:
	@for f in ${FOLDERS}; do 		\
    	cd $$f && make $@; 			\
		cd ${BASEPATH};				\
		echo "${GREEN}${BOLD}$$f${S}${S} has been built 🎉";	\
	done

re:
	@for f in ${FOLDERS}; do 			\
    	cd $$f && make $@; 				\
		cd ${BASEPATH};					\
		echo "${GREEN}${BOLD}$$f${S}${S} has been rebuilt 🎉";	\
	done

solutions:
	@for f in ${FOLDERS}; do 				\
    	cd $$f/${SOLUTIONS} && make re; 	\
		cd ${BASEPATH}/..;					\
		echo "${GREEN}${BOLD}$$f solution${S}${S} has been rebuilt 🎉";	\
	done

.PHONY: all re solutions

S 		=		\033[0m
BOLD 	= 		\033[1m
ITALIC 	= 		\033[3m
UNDER 	= 		\033[4m
REV 	= 		\033[7m

# Colors
GREY 	= 		\033[30m
RED 	= 		\033[31m
GREEN	=		\033[32m
YELLOW	=		\033[33m
BLUE	=		\033[34m
PURPLE	=		\033[35m
CYAN	=		\033[36m
WHITE	=		\033[37m

SGREY	=		\033[90m
SRED	=		\033[91m
SGREEN	=		\033[92m
SYELLOW	=		\033[93m
SBLUE	=		\033[94m
SPURPLE	=		\033[95m
SCYAN	=		\033[96m
SWHITE	=		\033[97m

# Colored backgrounds

IGREY	=		\033[40m
IRED	=		\033[41m
IGREEN	=		\033[42m
IYELLOW	=		\033[43m
IBLUE	=		\033[44m
IPURPLE	=		\033[45m
ICYAN	=		\033[46m
IWHITE	=		\033[47m