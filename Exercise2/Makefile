# You have to put your .c files in the FILES variable.
FILES_1 =	helloworld.c
FILES_2	=	max3.c exerciseTwo_questionOne.c
FILES_3 = 	max3.c max3_argv.c
FILES_4	=	max3.c max3_stdin.c
FILES_5	=	dayofbirth.c exerciseThree_questionOne.c
FILES_6	=	dayofbirth.c dayofbirth_argv.c
FILES_7	=	dayofbirth.c dayofbirth_stdin.c
FILES_8 =	pointers.c exerciseFour_questionOne.c
FILES_9	=	pointers.c exerciseFour_questionTwo.c
FILES_10 =	string.c exerciseFive_questionOne.c
FILES_11 =	string.c exerciseFive_questionTwo.c

S_PATH	=	srcs/
O_PATH	=	objs/
I_PATH	=	includes/
B_PATH	=	build/

SRCS_1	=	${addprefix ${S_PATH}, ${FILES_1}}
OBJS_1	=	${addprefix ${O_PATH}, ${FILES_1:.c=.o}}

SRCS_2	=	${addprefix ${S_PATH}, ${FILES_2}}
OBJS_2	=	${addprefix ${O_PATH}, ${FILES_2:.c=.o}}

SRCS_3	=	${addprefix ${S_PATH}, ${FILES_3}}
OBJS_3	=	${addprefix ${O_PATH}, ${FILES_3:.c=.o}}

SRCS_4	=	${addprefix ${S_PATH}, ${FILES_4}}
OBJS_4	=	${addprefix ${O_PATH}, ${FILES_4:.c=.o}}

SRCS_5	=	${addprefix ${S_PATH}, ${FILES_5}}
OBJS_5	=	${addprefix ${O_PATH}, ${FILES_5:.c=.o}}

SRCS_6	=	${addprefix ${S_PATH}, ${FILES_6}}
OBJS_6	=	${addprefix ${O_PATH}, ${FILES_6:.c=.o}}

SRCS_7	=	${addprefix ${S_PATH}, ${FILES_7}}
OBJS_7	=	${addprefix ${O_PATH}, ${FILES_7:.c=.o}}

SRCS_8	=	${addprefix ${S_PATH}, ${FILES_8}}
OBJS_8	=	${addprefix ${O_PATH}, ${FILES_8:.c=.o}}

SRCS_9	=	${addprefix ${S_PATH}, ${FILES_9}}
OBJS_9	=	${addprefix ${O_PATH}, ${FILES_9:.c=.o}}

SRCS_10	=	${addprefix ${S_PATH}, ${FILES_10}}
OBJS_10	=	${addprefix ${O_PATH}, ${FILES_10:.c=.o}}

SRCS_11	=	${addprefix ${S_PATH}, ${FILES_11}}
OBJS_11	=	${addprefix ${O_PATH}, ${FILES_11:.c=.o}}

CC		=	clang

RM		=	rm -rf

CFLAGS		=	-DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls

${O_PATH}%.o:	${S_PATH}%.c
				@mkdir -p ${dir $@}
				@${CC} ${CFLAGS} -c $< -o $@
				@echo "${BOLD}${CYAN}Compiling${S}${S} ${IGREY}$<${S} ${YELLOW}-> ${S}${SBLUE}$@${S} ⚙️"


questionOneOne:		build ${OBJS_1}
					@${CC} ${OBJS_1} ${CFLAGS} -o $@ -I ${I_PATH}
					@echo ""
					@echo "${BOLD}${PURPLE}Building${S}${S} ${IGREY}$@${S} 🖥️"
					@echo ""
					@echo "\033[3;92mCompilation is completed !${S} 🎉"
					@mv ./$@ ${B_PATH}$@

questionTwoOne:		build ${OBJS_2}
					@${CC} ${OBJS_2} ${CFLAGS} -o $@ -I ${I_PATH}
					@echo ""
					@echo "${BOLD}${PURPLE}Building${S}${S} ${IGREY}$@${S} 🖥️"
					@echo ""
					@echo "\033[3;92mCompilation is completed !${S} 🎉"
					@mv ./$@ ${B_PATH}$@

questionTwoTwo:		build ${OBJS_3}
					@${CC} ${OBJS_3} ${CFLAGS} -o $@ -I ${I_PATH}
					@echo ""
					@echo "${BOLD}${PURPLE}Building${S}${S} ${IGREY}$@${S} 🖥️"
					@echo ""
					@echo "\033[3;92mCompilation is completed !${S} 🎉"
					@mv ./$@ ${B_PATH}$@

questionTwoThree:	build ${OBJS_4}
					@${CC} ${OBJS_4} ${CFLAGS} -o $@ -I ${I_PATH}
					@echo ""
					@echo "${BOLD}${PURPLE}Building${S}${S} ${IGREY}$@${S} 🖥️"
					@echo ""
					@echo "\033[3;92mCompilation is completed !${S} 🎉"
					@mv ./$@ ${B_PATH}$@

questionThreeOne:	build ${OBJS_5}
					@${CC} ${OBJS_5} ${CFLAGS} -o $@ -I ${I_PATH}
					@echo ""
					@echo "${BOLD}${PURPLE}Building${S}${S} ${IGREY}$@${S} 🖥️"
					@echo ""
					@echo "\033[3;92mCompilation is completed !${S} 🎉"
					@mv ./$@ ${B_PATH}$@

questionThreeTwo:	build ${OBJS_6}
					@${CC} ${OBJS_6} ${CFLAGS} -o $@ -I ${I_PATH}
					@echo ""
					@echo "${BOLD}${PURPLE}Building${S}${S} ${IGREY}$@${S} 🖥️"
					@echo ""
					@echo "\033[3;92mCompilation is completed !${S} 🎉"
					@mv ./$@ ${B_PATH}$@

questionThreeThree:	build ${OBJS_7}
					@${CC} ${OBJS_7} ${CFLAGS} -o $@ -I ${I_PATH}
					@echo ""
					@echo "${BOLD}${PURPLE}Building${S}${S} ${IGREY}$@${S} 🖥️"
					@echo ""
					@echo "\033[3;92mCompilation is completed !${S} 🎉"
					@mv ./$@ ${B_PATH}$@

questionFourOne:	build ${OBJS_8}
					@${CC} ${OBJS_8} ${CFLAGS} -o $@ -I ${I_PATH}
					@echo ""
					@echo "${BOLD}${PURPLE}Building${S}${S} ${IGREY}$@${S} 🖥️"
					@echo ""
					@echo "\033[3;92mCompilation is completed !${S} 🎉"
					@mv ./$@ ${B_PATH}$@

questionFourTwo:	build ${OBJS_9}
					@${CC} ${OBJS_9} ${CFLAGS} -o $@ -I ${I_PATH}
					@echo ""
					@echo "${BOLD}${PURPLE}Building${S}${S} ${IGREY}$@${S} 🖥️"
					@echo ""
					@echo "\033[3;92mCompilation is completed !${S} 🎉"
					@mv ./$@ ${B_PATH}$@

questionFiveOne:	build ${OBJS_10}
					@${CC} ${OBJS_10} ${CFLAGS} -o $@ -I ${I_PATH}
					@echo ""
					@echo "${BOLD}${PURPLE}Building${S}${S} ${IGREY}$@${S} 🖥️"
					@echo ""
					@echo "\033[3;92mCompilation is completed !${S} 🎉"
					@mv ./$@ ${B_PATH}$@

questionFiveTwo:	build ${OBJS_11}
					@${CC} ${OBJS_11} ${CFLAGS} -o $@ -I ${I_PATH}
					@echo ""
					@echo "${BOLD}${PURPLE}Building${S}${S} ${IGREY}$@${S} 🖥️"
					@echo ""
					@echo "\033[3;92mCompilation is completed !${S} 🎉"
					@mv ./$@ ${B_PATH}$@

all:		build questionOneOne questionTwoOne questionTwoTwo questionTwoThree questionThreeOne questionThreeTwo \
			questionThreeThree questionFourOne questionFourTwo questionFiveOne questionFiveTwo

clean:
			@${RM} ${O_PATH}*
			@echo "${BOLD}${SRED}Removing${S}${S} ${IGREY}${O_PATH}${S} 🗑️"

fclean:		clean
			@${RM} ${B_PATH}*
			@echo "${BOLD}${SRED}Removing${S}${S} ${IGREY}${B_PATH}${S} 🗑️"

space:
			@echo " "

re:			fclean space all

build: 		
			@mkdir -p $@

.PHONY: all clean fclean re build space

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