 SELECT A.EMPLID,
B.NAME_DISPLAY,
C.CLASS_NBR,
D.CATALOG_NBR,
E.COURSE_TITLE_LONG,
TO_CHAR(CAST((F.MEETING_TIME_START) AS TIMESTAMP),'HH24.MI.SS.FF'), 
TO_CHAR(CAST((F.MEETING_TIME_END) AS TIMESTAMP),'HH24.MI.SS.FF'), 
F.MON, F.TUES, F.WED, F.THURS, F.FRI, F.SAT, F.SUN, 
TO_CHAR(F.START_DT,'YYYY-MM-DD'), 
TO_CHAR(F.END_DT,'YYYY-MM-DD')
  FROM SYSADM.PS_ACAD_PROG A, SYSADM.PS_NAMES B, SYSADM.PS_STDNT_ENRL C, SYSADM.PS_CLASS_TBL D, SYSADM.PS_CRSE_CATALOG E, SYSADM.PS_CLASS_MTG_PAT F
  WHERE ( A.EFFDT =
        (SELECT MAX(A_ED.EFFDT) FROM SYSADM.PS_ACAD_PROG A_ED
        WHERE A.EMPLID = A_ED.EMPLID
          AND A.ACAD_CAREER = A_ED.ACAD_CAREER
          AND A.STDNT_CAR_NBR = A_ED.STDNT_CAR_NBR
          AND A_ED.EFFDT <= SYSDATE)
    AND A.EFFSEQ =
        (SELECT MAX(A_ES.EFFSEQ) FROM SYSADM.PS_ACAD_PROG A_ES
        WHERE A.EMPLID = A_ES.EMPLID
          AND A.ACAD_CAREER = A_ES.ACAD_CAREER
          AND A.STDNT_CAR_NBR = A_ES.STDNT_CAR_NBR
          AND A.EFFDT = A_ES.EFFDT)
     AND A.PROG_STATUS = 'AC'
     AND A.EMPLID = B.EMPLID
     AND B.EFFDT =
        (SELECT MAX(B_ED.EFFDT) FROM SYSADM.PS_NAMES B_ED
        WHERE B.EMPLID = B_ED.EMPLID
          AND B.NAME_TYPE = B_ED.NAME_TYPE
          AND B_ED.EFFDT <= SYSDATE)
     AND B.NAME_TYPE = 'PRI'
     AND B.EFF_STATUS = 'A'
     AND A.EMPLID = C.EMPLID
     AND C.INSTITUTION = A.INSTITUTION
     AND C.ACAD_CAREER = D.ACAD_CAREER
     AND C.INSTITUTION = D.INSTITUTION
     AND C.STRM = D.STRM
     AND C.CLASS_NBR = D.CLASS_NBR
     AND D.SESSION_CODE = C.SESSION_CODE
     AND C.STDNT_ENRL_STATUS = 'E'
     AND D.CRSE_ID = E.CRSE_ID
     AND E.EFFDT =
        (SELECT MAX(E_ED.EFFDT) FROM SYSADM.PS_CRSE_CATALOG E_ED
        WHERE E.CRSE_ID = E_ED.CRSE_ID
          AND E_ED.EFFDT <= SYSDATE)
     AND D.CRSE_ID = F.CRSE_ID
     AND D.CRSE_OFFER_NBR = F.CRSE_OFFER_NBR
     AND D.STRM = F.STRM
     AND D.SESSION_CODE = F.SESSION_CODE
     AND D.CLASS_SECTION = F.CLASS_SECTION
     AND D.STRM = '1840'
     AND A.EMPLID = '0000000025')