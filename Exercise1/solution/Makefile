FILES =		main.c			\
			calculator.c	\
			maths.c			\


S_PATH	=	srcs/
O_PATH	=	objs/
I_PATH	=	includes/
B_PATH	=	build/

SRCS	=	${addprefix ${S_PATH}, ${FILES}}
OBJS	=	${addprefix ${O_PATH}, ${FILES:.c=.o}}

NAME	=	calculator_and_math

CC		=	clang

RM		=	rm -rf

CFLAGS	=	-DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls

${O_PATH}%.o:	${S_PATH}%.c
				@mkdir -p ${dir $@}
				@${CC} ${CFLAGS} -c $< -o $@ -I @{I_PATH}
				@echo "${BOLD}${CYAN}Compiling${S}${S} ${IGREY}$<${S} ${YELLOW}-> ${S}${SBLUE}$@${S} ⚙️"


${NAME}:	${OBJS}
			@${CC} ${OBJS} ${CFLAGS} -o $@ -I ${I_PATH}
			@echo ""
			@echo "${BOLD}${PURPLE}Building${S}${S} ${IGREY}$@${S} 🖥️"
			@echo ""
			@echo "\033[3;92mCompilation is completed !${S} 🎉"
			@mv ./$@ ${B_PATH}$@

all:		build ${NAME}

clean:
			@${RM} ${O_PATH}*
			@echo "${BOLD}${SRED}Removing${S}${S} ${IGREY}${O_PATH}${S} 🗑️"

fclean:		clean
			@${RM} ${B_PATH}*
			@echo "${BOLD}${SRED}Removing${S}${S} ${IGREY}${B_PATH}${S} 🗑️"

space:
			@echo " "

re:			fclean space all

test:		re
			@echo ""
			@echo "${BOLD}${PURPLE}Launching tests${S}${S} 🖥️"
			@echo ""
			@${B_PATH}${NAME}

build: 		
			@mkdir -p $@

.PHONY: all clean fclean re build space test

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